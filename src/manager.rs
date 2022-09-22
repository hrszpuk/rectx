/**  Manager module
 *   This module manages a rectx project!
 *   Features:
 *   - Generating new projects
 *   - Building projects
 *   - Running projects
**/

use std::fs;
use std::io::Write;
use std::path::Path;
use std::process::Command;
use fs_extra::dir::CopyOptions;
use crate::cli;
use crate::config::{Config, Profile};

/// Generates a project directory containing the following:
/// 1. README.md (with project name)
/// 2. /src/main.rct (with hello world program)
pub fn generate_project_directory(name: &String) -> std::io::Result<()> {

    // We generate a few directories: {name}/, {name}/src/, and {name}/deps/
    fs::create_dir_all(format!("{}/src", name))?;
    fs::create_dir(format!("{}/deps", name))?;

    // Here we create main.rct and README.md files
    // We also write basic text to these files:
    // - main.rct: hello world program
    // - README.md: project title
    let mut main = fs::File::create(format!("{}/src/main.rct", name))?;
    main.write_all(b"package sys;\n\nsys::Print(\"Hello, World!\");\n")?;

    let mut readme = fs::File::create(format!("{}/README.md", name))?;
    readme.write_all(format!("# {}\n", name).as_ref())?;

    // Generate a new config file
    let conf = Config::new(name);
    conf.generate(name)?;

    Ok(())
}

/// Uses rgoc to create an executable file from /src/main.rct
pub fn generate_project_executable(run: bool) -> std::io::Result<()> {

    // Load config file
    let conf = Config::load("config.toml");
    let profile: Profile;
    let profile_name: String;
    if run {
        profile = conf.run;
        profile_name = String::from("run");
    } else {
        profile = conf.build;
        profile_name = String::from("build");
    };

    let build_path = format!("{}/{}", &conf.project.target, &profile_name);

    let main = format!("{}/{}", &profile.source_dir, &profile.source_main);

    if !Path::new(&build_path).exists() {
        fs::create_dir_all(&build_path)?;
    }

    let dir_paths = fs::read_dir(&profile.source_dir)?;

    // Functional control flow magic
    let paths = dir_paths.filter_map(|entry| {
        entry.ok().and_then(|e|
            e.path().file_name()
                .and_then(|n| n.to_str().map(|s| String::from(s)))
        )
    }).collect::<Vec<String>>();

    // Building file
    if paths.contains(&profile.source_main) {
        let mut child = Command::new(&profile.compiler);
        for flag in profile.compiler_flags {
            child.arg(&flag);
        }
        child
            .arg(format!("-o={}/{}/{}", &conf.project.target, &profile_name, &profile.output_name))
            .arg(&main)
            .spawn()?
            .wait()?;
    } else {
        cli::abort(format!(
            "Failed to find target \"{}\" in source directory \"{}\"!",
            &profile.source_main,
            &profile.source_dir,
        ))
    }

    // Moving dependencies into target
    match fs_extra::dir::copy(
        &conf.project.dependencies,
        format!("{}/{}", &conf.project.target, &profile_name),
        &CopyOptions::new(),
    ) {
        Ok(_) => println!("Moved dependencies into target successfully!"),
        Err(_) => cli::abort("Unable to move dependencies!".to_string()),
    }

    if run {
        cli::process("Running project executable".to_string());

        let mut child = Command::new(
            format!("./{}/{}/{}", &conf.project.target, &profile_name, &profile.output_name)
        )
            .spawn()?;
        child.wait()?;
    }

    Ok(())
}
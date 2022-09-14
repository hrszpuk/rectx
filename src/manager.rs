/**  Manager module
 *   This module manages a rectx project!
 *   Features:
 *   - Generating new projects
 *   - Building projects
 *   - Running projects
**/

use std::{env, fs};
use std::io::Write;
use std::process::Command;
use crate::cli;
use crate::config;

/// Generates a project directory containing the following:
/// 1. README.md (with project name)
/// 2. /src/main.rct (with hello world program)
pub fn generate_project_directory(name: &String) -> std::io::Result<()> {

    // We generate a few directories: /{name} and /{name}/src
    fs::create_dir_all(name)?;
    fs::create_dir(format!("{}/src", name))?;

    // Here we create main.rct and README.md files
    // We also write basic text to these files:
    // - main.rct: hello world program
    // - README.md: project title
    let mut main = fs::File::create(format!("{}/src/main.rct", name))?;
    main.write_all(b"package sys;\n\nsys::Print(\"Hello, World!\");\n")?;

    let mut readme = fs::File::create(format!("{}/README.md", name))?;
    readme.write_all(format!("# {}\n", name).as_ref())?;

    let mut conf = config::Config::new(name);
    conf.generate(name).expect("TODO: panic message");

    Ok(())
}

/// Uses rgoc to create an executable file from /src/main.rct
pub fn generate_project_executable() -> std::io::Result<()> {

    // Getting source files
    let dir_paths = fs::read_dir("./src")?;
    let mut paths = Vec::new();

    for path in dir_paths {

        // TODO: check if file is rect file (.rct)
        let x = path?.file_name();
        let x = match x.into_string() {
            Ok(filename) => filename,
            Err(error) => {
                println!("{:?}", error);
                String::new()
            }
        };
        paths.push(x);
    }

    // Building file
    if paths.contains(&String::from("main.rct")) {
        let mut child = Command::new("rgoc")
            .arg("./src/main.rct")
            .spawn()?;
        child.wait()?;
    } else {
        println!("rectx :: Could not find \"main.rct\" in \"/src\"!");
    }

    Ok(())
}

/// Generates an executable file using generate_project_executable()
/// and then runs the executable that's generated.
pub fn generate_executable_and_run() -> std::io::Result<()>{
    generate_project_executable()?;

    cli::process(String::from("Running project executable"));

    let mut child = Command::new("./src/main")
        .spawn()?;
    child.wait()?;
    Ok(())
}
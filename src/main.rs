use std::env;

fn main() {
    cli::process_flags(env::args());
}

/// Manages Command Line Interface for getting commands, subcommands, and flags.
mod cli {
    use std::process::exit;
    use crate::manager;
    use crate::manager::generate_project;

    /// Takes the command line arguments and calls the correct function.
    pub fn process_flags(args: std::env::Args) {
        let mut arguments = Vec::new();
        for arg in args {
            arguments.push(arg);
        }
        arguments.remove(0);  // Drop application name ./rectx


        for arg in &arguments {
            match arg.as_str() {
                "help" => help(),
                "new" => new_project(&arguments),
                "build" => build_project(),
                "run" => run_project(),
                _ => help()
            }
        }
    }

    /// Show help menu to the user.
    /// The help menu contains information on commands and flags, and what they do.
    pub fn help() {
        println!(
            ":: ReCTx Help Menu ::
ReCTx is a project manager for the ReCT programming language!

Create a new project:
$ rectx new project-name

Run the project:
$ rectx run

Run the project:
$ rectx build
"
        );
        exit(0);
    }

    /// Creates a new project.
    /// This is called when the user calls the command "new".
    pub fn new_project(args: &Vec<String>) {
        if args.len() > 1 {
            match manager::generate_project(&args[1]) {
                Ok(()) => println!("Created new project \"{}\"!", args[1]),
                Err(error) => {
                    println!("Failed to create project \"{}\"!", args[1]);
                    println!("{:?}", error);
                },
            };
        }
        exit(0);
    }

    /// Builds the project in the current directory.
    /// This is called when the user calls the command "build".
    pub fn build_project() {

    }

    /// Runs the project in the current directory.
    /// This is called when the user calls the command "run".
    pub fn run_project() {

    }
}

// Manages generating projects, build, and running
mod manager {
    use std::{env, fs};
    use std::io::Write;
    use std::process::Command;

    pub fn generate_project(name: &String) -> std::io::Result<()> {
        fs::create_dir_all(name)?;
        fs::create_dir(format!("{}/src", name))?;
        let mut main = fs::File::create(
            format!("{}/src/main.rct", name)
        )?;
        main.write_all(b"package sys;\n\nsys::Print(\"Hello, World!\");\n")?;
        let mut readme = fs::File::create(
            format!("{}/README.md", name)
        )?;
        readme.write_all(format!("# {}\n", name).as_ref())?;
        Command::new("git")
            .arg("init")
            .arg(name)
            .spawn()?;

        Ok(())
    }
}

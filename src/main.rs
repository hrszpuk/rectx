use std::env;

fn main() {
    cli::process_flags(env::args());
}

/// Manages Command Line Interface for getting commands, subcommands, and flags.
mod cli {
    use std::process::exit;
    use crate::manager;

    /// Takes the command line arguments and calls the correct function.
    pub fn process_flags(args: std::env::Args) {
        let mut arguments = Vec::new();
        for arg in args {
            arguments.push(arg);
        }
        arguments.remove(0);  // Drop application name ./rectx


        if arguments.len() > 0 {
            match arguments[0].as_str() {
                "help" => help(),
                "new" => if arguments.len() > 1 {
                    new_project(&arguments)
                } else {
                    help_new()
                },
                "build" => build_project(),
                "run" => run_project(),
                _ => help_unknown()
            }
        } else {
            help()
        }
    }

    /// Show help menu to the user.
    /// The help menu contains information on commands and flags, and what they do.
    pub fn help() {
        println!(
            "ReCTx :: Help Menu

Usage: rectx <command> [options]

Commands:
  help          -> shows this menu
  new [name]    -> creates a new project
  run           -> runs the current project
  build         -> builds the current project

For more information visit the GitHub page: https://github.com/hrszpuk/rectx"
        );
        exit(0);
    }

    /// A specific help menu for the "new" command
    /// Creating a new project using the new command: usage and explanation
    pub fn help_new() {
        println!(
            "ReCTx :: Help Menu :: \"new\"

Usage: rectx new project-name

This command will create a new ReCT project with the name provided.
The project will contain: /src/main.rs, README.md, and config.toml!

For more information visit the GitHub page: https://github.com/hrszpuk/rectx"
        )
    }

    pub fn help_unknown() {
        println!(
            "ReCTX :: Help Menu :: Unknown

Usage: rectx <command> [options]

The command you have entered does not seem to exist!
Use \"rectx help\" for more information on the commands you can use.

For more information visit the GitHub page: https://github.com/hrszpuk/rectx"
        )
    }

    /// Creates a new project.
    /// This is called when the user calls the command "new".
    pub fn new_project(args: &Vec<String>) {
        if args.len() > 1 {
            match manager::generate_project_directory(&args[1]) {
                Ok(()) => println!("rectx :: Created new project \"{}\"!", args[1]),
                Err(error) => {
                    println!("rectx :: Failed to create project \"{}\"!", args[1]);
                    println!("{:?}", error);
                },
            };
        }
        exit(0);
    }

    /// Builds the project in the current directory.
    /// This is called when the user calls the command "build".
    pub fn build_project() {
        println!("Building project!");
        match manager::generate_project_executable() {
            Ok(()) => println!("rectx :: Executable was generated successfully!"),
            Err(error) => {
                println!("rectx :: Failed to generate executable!");
                println!("{:?}", error);
            }
        }
        exit(0);
    }

    /// Runs the project in the current directory.
    /// This is called when the user calls the command "run".
    pub fn run_project() {
        match manager::generate_executable_and_run() {
            Ok(()) => println!("rectx :: Executable was generated successfully!"),
            Err(error) => {
                println!("rectx :: Execution failed!")
            }
        }
        exit(0);
    }
}

// Manages generating projects, build, and running
mod manager {
    use std::{env, fs};
    use std::io::Write;
    use std::process::Command;

    /// Generates a project directory containing the following:
    /// 1. README.md (with project name)
    /// 2. /src/main.rct (with hello world program)
    pub fn generate_project_directory(name: &String) -> std::io::Result<()> {
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

        Ok(())
    }

    /// Uses rgoc to create an executable file from /src/main.rct
    pub fn generate_project_executable() -> std::io::Result<()> {
        // Getting source files
        let dir_paths = fs::read_dir("./src")?;
        let mut paths = Vec::new();
        for path in dir_paths {
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
            Command::new("rgoc")
                .arg("./src/main.rct")
                .spawn()?;
        } else {
            println!("rectx :: Could not find \"main.rct\" in \"/src\"!");
        }

        Ok(())
    }

    /// Generates an executable file using generate_project_executable()
    /// and then runs the executable that's generated.
    pub fn generate_executable_and_run() -> std::io::Result<()>{
        generate_project_executable()?;
        Command::new("./src/main")
            .spawn()?;
        Ok(())
    }
}


mod config {
    use std::fs;
    use std::process::exit;
    use serde_derive::{Deserialize, Serialize};


    #[derive(Debug, Deserialize, Serialize)]
    struct Config {
        pub project_name: String,
    }

    impl Config {
        pub fn new(name: String) -> Config {
            Config {
                project_name: name,
            }
        }

        pub fn generate(&self, path: String) -> std::io::Result<()> {
            let toml_data = match toml::to_string(&self) {
                Ok(data) => data,
                Err(error) => {  // Temporary hacky fix -- definitely change later
                    println!("ERROR! {:?}", error);
                    exit(1);
                    String::new()
                }
            };
            fs::write(path + "/config.toml", toml_data)?;
            Ok(())
        }

        /*pub fn load(path: String) -> std::io::Result<Config> {
            let data = fs::read(path)?;
            toml::from_slice(&*data)?
        }*/
    }
}

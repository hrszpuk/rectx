/**  Cli module
*   This module manages the command line interface of rectx!
*   Features:
*   - Handling program arguments (commands and options)
*   - Displaying help menus
*   - Recognising commands and calling the correct functions
 **/

use std::process::exit;
use crate::manager;

/// Takes the command line arguments and calls the correct function.
pub fn process_flags(args: std::env::Args) {

    // Moving the program arguments into a custom vector.
    // We do this so we can remove some parts of the program arguments
    // such as the application name.
    let mut arguments = Vec::new();
    for arg in args {
        arguments.push(arg);
    }

    // Drop application name ./rectx
    arguments.remove(0);

    // Here we just match for different commands
    // and call the correct function.
    // Example: "new" -> new_project()
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
        // Here we call the help menu if we only encounter "rectx"
        help()
    }
}

/// Creates a new project.
/// This is called when the user calls the command "new".
pub fn new_project(args: &Vec<String>) {
    if args.len() > 1 {

        // Checking for errors
        match manager::generate_project_directory(&args[1]) {

            // TODO: Add a cli::success or smth to handle outputs
            Ok(()) => println!("rectx :: Created new project \"{}\"!", args[1]),
            Err(error) => {

                // TODO: Make error more explicit
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

    // TODO: Add more status messages to other functions (like below)
    println!("Building project!");

    // Checking for errors
    match manager::generate_project_executable() {

        // TODO: Success message here too...
        Ok(()) => println!("rectx :: Executable was generated successfully!"),
        Err(error) => {

            // TODO: Better error handling (multi-types)
            // TODO: Add ISSUE and ERROR cli handles
            println!("rectx :: Failed to generate executable!");
            println!("{:?}", error);
        }
    }
    exit(0);
}

/// Runs the project in the current directory.
/// This is called when the user calls the command "run".
pub fn run_project() {

    // Checking for errors
    match manager::generate_executable_and_run() {

        // TODO: Again, CLI handles here, and better error handling
        Ok(()) => println!("rectx :: Executable was generated successfully!"),
        Err(error) => {
            println!("rectx :: Execution failed!")
        }
    }
    exit(0);
}

/// Use cli::success when a successful process has taken place
pub fn success(message: &String) {
    println!(
        "[SUCCESS] {}", message
    );
}

/// Use cli::process when a process has began (logging)
pub fn process(message: &string) {
    println!(
        "::: {}", message
    );
}

/// use cli::issue when an issue is found, but can be recovered
pub fn issue(message: &String) {
    println!(
        ":!: {}", message
    );
}

/// Use cli::abort when an unrecoverable issue is found (exits the program)
pub fn abort(message: &String) {
    println!(
        "!!! {}", message
    );
    println!(
        "[ABORT] An unrecoverable error caused rectx to abort!"
    );
    exit(2);
}

/// Use cli::info when wanting to give the user non-specific information
pub fn info(message: &String) {
    println!(
        "[INFO] {}", message
    );
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

/// A specific help menu for the "new" command.
/// Creating a new project using the new command: usage and explanation.
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
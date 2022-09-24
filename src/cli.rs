/**  Cli module
*   This module manages the command line interface of rectx!
*   Features:
*   - Handling program arguments (commands and options)
*   - Displaying help menus
*   - Recognising commands and calling the correct functions
 **/
use std::{fs, io};
use std::io::{ErrorKind, stdout, Write};
use std::process::exit;
use colored::Colorize;
use crate::manager;
use crate::manager::generate_project_directory;

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
            "help" => {
                if arguments.len() > 1 {
                    match arguments[1].as_str() {
                        "help" => help(),
                        "new"|"create" => help_new(),
                        "build" => help_build(),
                        "run" => help_run(),
                        _ => help(),
                    }
                } else {
                    help();
                }
            },
            "new"|"create" => new_project(&arguments),
            "build" => build_project(false),
            "run" => build_project(true),
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

        process(format!("Creating a new project named \"{}\"", args[1]));

        // Checking for errors
        match generate_project_directory(&args[1]) {

            Ok(()) => success(format!("Created new project \"{}\"", args[1])),
            Err(error) => match error.kind() {

                // Will occur if "main.rs" or "readme" fail to be created and
                // the write_all method is still triggered. This is unlikely,
                // but is handled anyway (just in case ofc).
                ErrorKind::NotFound => abort(
                    String::from("A previously generated file could not be found!")
                ),

                // Will occur if rectx doesn't have permission to generate project files or folders
                // where the user requests it to be generated.
                // A common cause of this could be user attempting to create a new project in
                // administrator only directories (such as / on *nix).
                ErrorKind::PermissionDenied => abort(
                    String::from(
                        "Attempted to generate project entities but received permission denied!"
                    )
                ),

                // Will occur if a directory of this name already exists.
                // Here we ask if the user would like to overwrite the directory,
                // and if not we simply abort
                ErrorKind::AlreadyExists => {
                    issue(format!(
                        "The directory \"{}\" already exists and therefore could not be generated!",
                        args[1],
                    ));

                    // Here we ask about overwriting the directory
                    let overwrite = question(format!("Would you like to overwrite the directory \"{}\"?", args[1]));
                    if overwrite {

                        // Of course overwriting could also error, in that case we abort
                        process(String::from("Attempting to remove directory!"));
                        match fs::remove_dir_all(&args[1]) {

                            // Overwriting successful, we retry to create project
                            Ok(()) => {
                                success(format!("Removed directory \"{}\"!", args[1]));
                                process(String::from("Retrying project creation..."));
                                generate_project_directory(&args[1]).expect("!!! Project creation failed again... aborting!");
                                success(format!("Created new project \"{}\"", args[1]))
                            },

                            // Overwriting unsuccessful, quickly abort
                            Err(_) => abort(
                                String::from("Attempt to overwrite directory failed!")
                            )
                        }
                    } else {

                        // This is if the user doesn't want to overwrite btw
                        abort(
                            String::from("Failed to generate project because directory already exists!")
                        );
                    }
                },
                // Will occur if no more memory is avaiable on the user's machine
                // as we allocate memory in generate_project_directory
                ErrorKind::OutOfMemory => abort(
                    format!(
                        "Could not generate project files for \"{}\" because machine is out of memory!",
                        args[1],
                    )
                ),
                _ => abort(
                    String::from(
                        "Unable to recover from an unexpected error while generating new project!",
                    )
                )
            }
        };
    } else {
        help_new();
    }
    exit(0);
}

/// Builds the project in the current directory.
/// This is called when the user calls the command "build".
/// This also handles running the executable with argument "run"
pub fn build_project(run: bool) {

    process(String::from("Building project executable"));

    // Checking for errors
    match manager::generate_project_executable(run) {

        Ok(()) => success(String::from("Executable was successful!")),
        Err(error) => match error.kind() {
            ErrorKind::NotFound => abort(
                String::from("Could not generate executable as source file(s) could not be found!")
            ),
            ErrorKind::PermissionDenied => abort(
                String::from("Could not generate executable due to lack of permissions!")
            ),
            ErrorKind::OutOfMemory => abort(
                String::from("Could not generate executable as machine is out of memory!")
            ),
            _ => abort(
                String::from("Unexpected unrecoverable error occurred while generating project executable!")
            ),
        },
    }
    exit(0);
}

/// Use cli::success when a successful process has taken place
pub fn success(message: String) {
    println!(
        "[SUCCESS] {}", message
    );
}

/// Use cli::process when a process has began (logging)
pub fn process(message: String) {
    println!(
        "::: {}", message
    );
}

/// use cli::issue when an issue is found, but can be recovered
pub fn issue(message: String) {
    println!(
        ":!: {}", message
    );
}

/// Use cli::abort when an unrecoverable issue is found (exits the program)
pub fn abort(message: String) {
    println!(
        "!!! {}", message
    );
    println!(
        "[ABORT] An unrecoverable error caused rectx to abort!"
    );
    panic!();
}

/// Use cli::info when wanting to give the user non-specific information
#[allow(dead_code)]
pub fn info(message: String) {
    println!(
        "[INFO] {}", message
    );
}

/// Use cli::question when wanting to get information from the user
pub fn question(message: String) -> bool {
    let mut buffer = String::new();

    print!("{} [Y/n] ", message);
    stdout().flush().unwrap();

    match io::stdin().read_line(&mut buffer) {
        Ok(_) => (),
        Err(_) => abort(
            format!(
                "Attempted to question user but found unrecoverable error instead!"
            )
        )
    }

    match buffer.to_lowercase().trim() {
        "y"|"yeah"|"yes"|"ye"|"yeet"|"yee" => true,
        _ => false
    }
}

/// Show help menu to the user.
/// The help menu contains information on commands and flags, and what they do.
pub fn help() {
    println!(
        "{}{} {}\n\n{}",
        get_help_title(),
        "Usage:".bold(),
        "rectx <command>".bright_green(),
        "[Commands]".bold()
    );
    for (name, example, description) in get_help_commands() {
        println!(
            "{}\t {} {}\t {} {}",
            name.bold().bright_cyan(),
            ":".black(),
            example.bold().bright_blue(),
            ":".black(),
            description.bold().bright_green(),
        )
    }

    print_help_info();
    exit(0);
}

/// A specific help menu for the "new" command.
/// Creating a new project using the new command: usage and explanation.
pub fn help_new() {
    println!(
        "{}{} {}\n\n{}
The {} command will create a {}.
After calling this command you will be {} about the project such as:
- {}
- {}
- {}

An alias of {} is {} ({}) which functions the same.",
        get_help_title(),
        "Usage:".bold(),
        "rectx new".bright_green(),
        "[Description]".bold(),
        "new".bold().underline().cyan(),
        "new project in a new directory".bold().underline().magenta(),
        "prompted for information".bold().underline().blue(),
        "Author".blue().bold(),
        "Project name".blue().bold(),
        "Licensing options".blue().bold(),
        "new".cyan().bold().underline(),
        "create".cyan().bold().underline(),
        "rectx create".bright_green()
    );
    print_help_info()
}

pub fn help_build() {
    println!(
        "{}{} {}\n\n{}
The {} command can be used to {} from a ReCT project.
The {} command takes information from the {} (generated by {}),
and uses it to create the executable.

This means you can specify the {}!",
        get_help_title(),
        "Usage:".bold(),
        "rectx build".bright_green(),
        "[Description]".bold(),
        "build".cyan().bold().underline(),
        "create an executable".blue(),
        "build".cyan().bold().underline(),
        "config.toml".bold().underline(),
        "rectx new".bright_green(),
        "compiler, compiler flags, executable name, and more".bold().underline(),
    );
}

pub fn help_run() {
    println!(
        "{}{} {}\n\n{}
The {} command can be used to {} from a ReCT project.
The {} command takes information from the {} (generated by {}), and uses it to create the executable.

This means you can specify the {}!

The major difference between the {} command and the {} command is that the {} command {} for you.
{} the {} command has separate settings from the {} command in the {}.",
        get_help_title(),
        "Usage:".bold(),
        "rectx run".bright_green(),
        "[Description]".bold(),
        "run".cyan().bold().underline(),
        "create an executable".blue(),
        "run".cyan().bold().underline(),
        "config.toml".bold().underline(),
        "rectx new".bright_green(),
        "compiler, compiler flags, executable name, and more".bold().underline(),
        "run".cyan().bold().underline(),
        "build".cyan().bold().underline(),
        "run".cyan().bold().underline(),
        "executes the executable".bold().underline(),
        "IMPORTANT NOTE:".magenta(),
        "run".cyan().bold().underline(),
        "build".cyan().bold().underline(),
        "config.toml".bold().underline(),
    );
}

pub fn help_unknown() {
    println!(
        "{}{} {}
The command you have entered does not exist!
Use {} for information on the commands you can use.",
    get_help_title(),
    "Usage:".bold(),
    "rectx <command>".bright_green(),
    "rectx help".bold().underline().cyan()
    );
    print_help_info();
}

// Helper commands for printing help menu
// print_help_info() displays the github page link and rect discord server link
// get_help_commands() returns a vector of each help command
// get_help_title() returns a string of the help menu title

fn print_help_info() {
    println!("\n{}{}",
             "For more information check out the GitHub page: ".bold(),
             "https://github.com/hrszpuk/rectx".bright_blue());
    println!("{} {}",
             "or join the ReCT Discord server:".bold(),
             "https://discord.gg/Ymm9xGxWZf".bright_blue());
}

fn get_help_commands() -> Vec<(&'static str, &'static str, &'static str)> {
    vec![
        ("Help", "rectx help", "Shows this help message."),
        ("New", "rectx new", "Creates a new ReCTx project."),
        ("Build", "rectx build", "Builds the current ReCTx project."),
        ("Run", "rectx run", "Runs the current ReCTx project.")
    ]
}

fn get_help_title() -> String {
    let title = "ReCTx Project Manager";
    let version = "v1.0.0";
    let mut dashes = String::new();
    for _ in 0..(title.len()+version.len()+2) {
        dashes.push('-');
    }
    format!("{}\n{} {}\n{}\n", dashes, title.bold(), version.bright_green().bold(), dashes)
}
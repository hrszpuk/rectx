/**  Config module
*   This module manages the config.toml inside of a ReCTx project!
*   Features:
*   - Generating new config.toml files (for the manager module)
*   - Reading config.toml files into the Config struct
*   - Processing any issues regarding the config.toml (such as missing or incorrect fields)
 **/

use std::fs;
use std::process::exit;
use serde_derive::{Deserialize, Serialize};
use crate::cli;


/// Config structure will store all the contents within
/// the config.toml so Manager can use it turning compilation.
#[derive(Debug, Deserialize, Serialize)]
pub struct Config {
    pub project: Project,
    #[serde(rename(serialize = "profile-build", deserialize = "profile-build"))]
    pub build: Profile,
    #[serde(rename(serialize = "profile-run", deserialize = "profile-run"))]
    pub run: Profile,
}

/// Stores information about the project
/// This information comes from the [project] table in config.toml
#[derive(Debug, Deserialize, Serialize)]
pub struct Project {
    pub name: String,
    pub version: String,
    pub authors: Vec<String>,
    pub remote: String,
    pub target: String,
    pub dependencies: String,
}

/// Stores information about compilation profiles
/// This information comes from [profile.build] and [profile.run] in config.toml
#[derive(Debug, Deserialize, Serialize)]
pub struct Profile {
    pub compiler: String,
    pub compiler_flags: Vec<String>,
    pub source_dir: String,
    pub source_main: String,
    pub output_name: String,
}

impl Config {

    /// Simple constructor for Config struct
    /// Generates default config values for a new binary project
    pub fn new(name: &String) -> Config {
        Config {
            project: Project {
                name: name.clone(),
                version: String::from("0.1.0"),
                authors: vec![],
                remote: String::from(""),
                target: "target".to_string(),
                dependencies: "deps".to_string()
            },
            build: Profile {
                compiler: String::from("rgoc"),
                compiler_flags: vec!["-xx".to_string(), "-O".to_string()],
                source_dir: String::from("src"),
                source_main: String::from("main.rct"),
                output_name: name.clone(),
            },
            run: Profile {
                compiler: String::from("rgoc"),
                compiler_flags: vec!["-xx".to_string()],
                source_dir: String::from("src"),
                source_main: String::from("main.rct"),
                output_name: name.clone(),
            }
        }
    }

    /// Generates a config.toml file from the Config struct
    pub fn generate(&self, path: &String) -> std::io::Result<()> {

        // Serialises the config struct into a toml string
        let toml_data = match toml::to_string(&self) {
            Ok(data) => data,
            Err(error) => {
                cli::abort(
                    "Failed to serialize Config data!".to_string()
                );
                panic!();  // This is never reached because cli::abort() handles panics
            },
        };

        // Finally, we create the config.toml with our config data
        fs::write(format!("{}/config.toml", path), toml_data)?;
        Ok(())
    }

    /// Reads the contents of a config.toml and converts it into a config struct
    pub fn load(path: &str) -> Config {

        // TODO: config module errors
        let mut contents = fs::read_to_string(path).unwrap();
        let data = toml::from_str(contents.as_str()).unwrap();
        data
    }
}
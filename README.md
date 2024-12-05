# storctl

`storctl` is a command-line tool for managing and interacting with the STOR object store. It provides essential functionality for managing objects and buckets with a simple, intuitive interface.

---

## Features

- **List buckets**: View available buckets.
- **Create buckets**: Add new storage buckets.
- **Remove buckets**: Delete empty buckets.
- **List objects**: Browse contents of a bucket.
- **Copy objects**: Transfer objects between locations.
- **Move objects**: Relocate objects across buckets or paths.
- **Remove objects**: Delete specific objects from buckets.

---

## Installation

### From Source
1. Ensure you have Go installed (version 1.23 or higher).
2. Clone the repository:
   ```bash
   git clone https://github.com/cfichtmueller/storctl.git
   ```
3. Build the binary using the Makefile:
   ```bash
   cd storctl
   make binary
   ```
4. Move the binary to a directory in your `PATH` (e.g., `/usr/local/bin`).

---

## Usage

`storctl` commands follow the syntax:
```bash
storctl <command> [arguments] [options]
```

### Commands and Examples

#### List buckets (`lb`)
```bash
storctl lb
```
Example:
```bash
storctl lb 
# Output:
# NAME           OBJECTS   SIZE
# mybucket       8         1405614
# another-bucket 3         3160277
```

#### Create a bucket (`mb`)
```bash
storctl mb <bucket>
```
Example:
```bash
storctl mb newbucket
```

#### Remove a bucket (`rb`)
```bash
storctl rb <bucket>
```
Example:
```bash
storctl rb oldbucket
```

#### List objects (`ls`)
```bash
storctl ls <bucket>
```
Example:
```bash
storctl ls mybucket
# Output:
# file1.txt
# file2.txt
```

#### Copy an object (`cp`)
```bash
storctl cp <source> <destination>
```
Example:
```bash
storctl cp mybucket/file1.txt mybucket-backup/file1.txt
```

#### Move an object (`mv`)
```bash
storctl mv <source> <destination>
```
Example:
```bash
storctl mv mybucket/file1.txt mybucket-archive/file1.txt
```

#### Remove an object (`rm`)
```bash
storctl rm <bucket>/<object>
```
Example:
```bash
storctl rm mybucket/file1.txt
```

#### Configuration Management Commands

- **Create a configuration context**:
  ```bash
  storctl config create-context <context-name>
  ```
- **Delete a configuration context**:
  ```bash
  storctl config delete-context <context-name>
  ```
- **List all configuration contexts**:
  ```bash
  storctl config get-contexts
  ```
- **Rename a configuration context**:
  ```bash
  storctl config rename-context <old-name> <new-name>
  ```
- **Set credentials in a context**:
  ```bash
  storctl config set-credentials <context-name> --access-key <key> --secret-key <key>
  ```
- **Use a specific configuration context**:
  ```bash
  storctl config use-context <context-name>
  ```

### Initial Configuration

When you first run `storctl`, an empty configuration file is created automatically in your home directory. Use the `config` commands above to manage your configuration contexts.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


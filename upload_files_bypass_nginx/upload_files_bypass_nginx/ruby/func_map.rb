functions = {
  "system" => -> (command) { system(command) }
}

functions["system"].call("id")

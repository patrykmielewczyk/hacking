const { exec } = require("child_process");

exec("id", (error, stdout, stderr) => {
  console.log(stdout);
});

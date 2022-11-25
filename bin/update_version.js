const fs = require('fs');

const args = process.argv.slice(2);

if (args.length != 3) {
    console.error("Usage: node update_version.js [file] [pattern] [new_version]")
    process.exitCode = 1;
    return;
}

const file = fs.readFileSync(args[0]);
const fileContent = file.toString();

const regex = new RegExp(args[1].replace(/\//g, "\\/"))
if (regex.test(fileContent)) {
    const currentVersion = fileContent.match(regex)[1];

    const result = fileContent.replace(currentVersion, args[2])
    
    fs.writeFileSync(args[0], result)

    return
}

console.error("Couldn't find the previous version")
process.exitCode = 1;

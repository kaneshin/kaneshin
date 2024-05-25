// input: commence
// gpt returns the response below by using the gpt prompt in the same folder
let response = `Definitions: 開始する

Meaning: To begin or to start something. (何かを始める、もしくは開始すること。)

Phonics: /kəˈmens/

Part of speech: verb

Synonyms: begin, start, initiate, launch, inaugurate

Examples:
- The ceremony will commence at 10 AM. (式は午前10時に開始されます。)
- They commenced their journey early in the morning. (彼らは朝早く旅を始めました。)
- Construction on the new building will commence next week. (新しい建物の建設は来週開始されます。)`;

// behavior like Zapier.
let input = { response };
var output = {};

let lines = input.response.split("\n");
var examples = []; // examples has multi-lines
for (var i = 0; i < lines.length; i++) {
  let line = lines[i];
  switch (line.substr(0, 2)) {
    case "De":
      output["definition"] = line.match(/Definitions: (.*)/)[1];
      break;
    case "Me":
      output["meaning"] = line.match(/Meaning: (.*)/)[1];
      break;
    case "Ph":
      output["phonics"] = line.match(/Phonics: (.*)/)[1];
      break;
    case "Pa":
      output["speech"] = line.match(/Part of speech: (.*)/)[1];
      break;
    case "Sy":
      output["synonyms"] = line.match(/Synonyms: (.*)/)[1];
      break;
    case "- ":
      examples.push(line);
      break;
  }
}
output["examples"] = examples.join("\n");
// to notify slack of the response with blockquote visually
output["slack"] = input.response.replace(/^/, "> ").replaceAll(/\n/g, "\n> ");

console.log(output);

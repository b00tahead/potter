#!/usr/bin/env node

const fs = require("fs");
const OpenAI = require("openai");
const { argv } = require("yargs");

const CONFIG_FILE = "config.json";

// Function to initialize ChatGPT with API Key
function initialize(apiKey) {
  console.log(`Initializing ChatGPT with API Key: ${apiKey}`);
  // Store API key securely (e.g., in environment variables or config file)
  const config = { apiKey };
  fs.writeFileSync(CONFIG_FILE, JSON.stringify(config, null, 2));
}

// Function to load API Key from config file
function loadApiKey() {
  try {
    const configData = fs.readFileSync(CONFIG_FILE);
    const config = JSON.parse(configData);
    return config.apiKey;
  } catch (error) {
    console.error("Error loading API key:", error.message);
    return null;
  }
}

// Function to ask question to ChatGPT
async function askQuestion(question) {
  const apiKey = loadApiKey();
  if (!apiKey) {
    console.error(
      "API Key not found. Please run 'potter init YOUR_API_KEY' to initialize.",
    );
    return;
  }

  const openai = new OpenAI({
    apiKey: apiKey,
  });

  console.log(`Asking question: ${question}`);

  try {
    const response = await openai.chat.completions.create({
      model: "gpt-3.5-turbo-1106",
      messages: [
        {
          role: "user",
          content: question,
        },
      ],
      temperature: 1,
      max_tokens: 256,
      top_p: 1,
      frequency_penalty: 0,
      presence_penalty: 0,
    });

    console.log("Response:", response.choices[0].message.content);
    return;
  } catch (error) {
    console.error("Error:", error.message);
  }
}

// Main CLI logic
const command = argv._[0];

switch (command) {
  case "init":
    const apiKey = argv._[1];
    initialize(apiKey);
    break;
  default:
    const question = argv._.join(" ");
    askQuestion(question); // Replace 'YOUR_API_KEY' with actual API key
    break;
}

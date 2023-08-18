# gowhisper
Simple CLI program written in GO to get audio or video subtitles

To get started simply head over to the [openai website](https://platform.openai.com/account/api-keys) and create an API key. <br>

Simply install the program, set your API key, maybe set an output directory and get started transcribing.

gowhisper sk -k <YOUR_API_KEY> <br>
gowhisper sp -dir <YOUR_OUTPUT_DIRECTORY> (absolute path) <br>
gowhisper t -f <your_file.mp3>

At some point I might add more stuff like setting the language, output format etc. but that has to wait for now.

The program also doesn't do status updates, I saw someone implemented that in JavaScript but I don't really want to mess with their API too much

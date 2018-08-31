# simple-quiz
Simple command line quiz game to practice using go.

To Build
---
Use <code>go build</code> to build

To Run
---
To run the quiz, use <code>./simple-quiz</code>

Optional flags
---
- csv - Input a filename for the questions and answers. Use the format question,answer per line
- limit - Number of seconds for a time limit to complete the quiz. Omit if you don't want a time limit

Example: <code>./simple-quiz -csv="questions.csv" -limit=10</code>

Based on the exercises on [Gophercises](https://gophercises.com/)
# Inkwell

Inkwell is a golang-based CLI tool that takes a input file and formats its string via custom tags and punctuation corrections

Its also my first project, and one im proud of
<h2> usage </h2>

You give it a txt file, and within that file it will correct the punctuation and output it to another file result.txt
moreover, it detects <strong> tags </strong> 

<h2> Tags </h2>
!Ensure spaces before and after the tag before execution
<ul>
<li> Every instance of (bin) replaces the word (that is a binary number) before it with the decimal version of the word. </li>

<li> Every instance of (hex) replaces the word (that is a hexadecimal number) before it with the decimal version of the word. </li>

  <li> Every instance of (up) converts the word before with the Uppercase version of it. </li>

  <li> Every instance of (low) converts the word before with the Lowercase version of it. </li>

  <li> Every instance of (cap) converts the word before with the capitalized version of it. </li>
</ul>
  
 
<h3> Moreover for up, low, cap, a iterator feature is added, so if a number appears next to it, like so: </h3>

```
(low, <number>)
```
<h3> it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. </h3>

# wilde #
![image](http://kickassfacts.com/wp-content/uploads/2013/10/OscarWildes.jpg)
## Preface ##

I've started to see tv series in english which is not my mother tongue. While I'm listening, I write all the unknown words in order to check them up later. After some days, I have on my desk a lot of piece of paper, with strange words written on them.
For some of them I remember the translation, for others I completely forgot it.

I *need* something that grant to me to store all the translations for the terms that I find while I'm reading english, while I'm watching films in english and so on.

I *need* to have a *guide* that tells me what I've looked for yesterday or one year ago.

## Mission ##

Create a specific tool that, for each term retrieve all the translations associate to it, and let me choose one of them to save it in my digital translation store.
After that, it will display to me the terms saved in a specific period of time and also some statistics about the most seen terms.

-------------------------------------------------------------------------------

For all these reason I've decided to realize **Wilde** a simple Golang application that grant to store on a MongoDB database all the information about
the terms that you look for.

**With wilde, you'll never miss a term**

## Requisities ##

This golang application has been written using the latest version of the Golang binaries that you can download from the [official website] (http://golang.org/). 
The latest version of MongoDB has been used too and you can download it from the [official website] (http://mongodb.org).

## Configuration ##
Here will be described the sequence of operations needed to run **wilde**.
1. Create a MongoDB database with two different collections called: *terms* and *terms_trans*;
2. Create a configuration file using the template file [wilde.json] (https://github.com/aleSuglia/wilde/blob/master/wilde.json);
3. Run the command: `go install` within the wilde directory;

After these steps you will be able to go to $GOPATH/bin and execute wilde following the usage information.

## BE CAREFUL: This project is in active development so don't expect that it's fully working! ##

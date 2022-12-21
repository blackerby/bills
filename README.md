# bills
Law Library of Congress Remote Metadata Program PDF text extraction command line tool.

## Introduction
This is intended to be a simple to use program to expedite your work as a Remote Metadata Intern on the Congressional Research Service bill summaries project. Given a path to a CRS bill summary PDF and a page number, if there are any bill summaries on that page, `bills` will create a CSV file containing bill summary metadata and a text file for each bill summary it could find on that page.

## Installation
The most straightforward way to get up and running with this program is to [download the latest release](https://github.com/blackerby/bills/releases/tag/v0.0.1-alpha) and save the program in the same directory (folder) where the PDF you are working with is located. `bills-macos` is for Macs with Intel processors and `bills.exe` is a 64-bit binary for Windows. I plan to add a binary for Apple Silicon in the near future.

## Usage
This is a command line tool, meaning you need to be or get comfortable with working on the command line for your operating system. You can find a good introduction to working on the command line at [Library Carpentry](https://librarycarpentry.org/lc-shell/). In the examples that follow, `$` is used to represent the command prompt. Don't type it in on your command line.

Start by navigating to the directory where your PDF is located. This is also where you should have saved the copy of the program you downloaded earlier. Next, type the name of the program, followed by the flag used to indicate the name of the file you will be extracting text and data from, and finally the flag indicating the page to process. **Note:** this should be the page number of the PDF file itself, _not_ the page number on the original printed document.

Here's an example:
```bash
$ ./bills -file 74_2.pdf -page 3
```
If the program can find any bill summaries, you should see a list of filenames after running the program before you see your command prompt come up again.

The output from running the above command looks like this:
```bash
$ ./bills -file 74_2.pdf -page 3
74_2_s3_19350104.txt
74_2_s5_19350104.txt
74_2_s11_19350104.txt
74_2_s17_19350104.txt
74_2_s18_19350104.txt
```

Now, if you list the contents of the directory you are working in, you will not only see the files listed above, but also a CSV file:
```bash
$ ls
74_2_003.csv
74_2_s3_19350104.txt
74_2_s5_19350104.txt
74_2_s11_19350104.txt
74_2_s17_19350104.txt
74_2_s18_19350104.txt
...
```

From here you can go about manually editing the names and contents of the text files and importing the data from the CSV file into your working spreadsheet.

If the program cannot find a bill summary on the page you gave it, it will extract the text of that page into a single text file for you to review or delete. For example:

```bash
$ ./bills -file 74_2.pdf -page 1
Header pattern not found, writing full page to 74_2_001.txt
```

## Caveats
This program is pretty naive. It will save you some time and effort you would otherwise be spending on copying and pasting and manually naming text files, but it's not perfect because 1) I'm not a professional programmer and 2) the text data we are working with is pretty inconsistent at times (but it is far better than manually transcribing!).

Common difficulties you will run into include:
- The dates the program puts into the generated text files names are just the dates of introduction of the bill. The program does not look beyond the header line of a summary, so you will often need to correct the date in the name of a generated text file.
- You may find that the number of generated text files is smaller than the number of summaries you see on the page you are working on. The missing summary is probably hidden in the text file generated for the bill summary that comes right before it on the original page. This happens because the pattern that matches the header line for a bill summary has not accounted for some corner case. If you encounter this issue, I would [love to hear about it](mailto:wtblackerby@crimson.ua.edu).

## Problems
If you encounter a problem with the software and you have a GitHub account, please use the [issue tracker](https://github.com/blackerby/bills/issues) to let me know about it. If you don't have a GitHub account, [consider getting one](https://github.com/signup) 😉 and feel free to [send me an email](mailto:wtblackerby@crimson.ua.edu) to start a conversation

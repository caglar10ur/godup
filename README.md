# godup

Silly duplicate file finder that I used to find duplicated photos/movies in my iPhoto library

## Usage

* All files:

```
[caglar@Zangetsu] ./dup /Users/caglar/Pictures/iPhoto\ Library/Masters/
Started to walk over /Users/caglar/Pictures/iPhoto Library/Masters and its sub-directories to find duplicated files

1) 2 items
    /Users/caglar/Pictures/iPhoto Library/Masters/2012/07/18/20120718-032030/IMG_0407.MOV
    /Users/caglar/Pictures/iPhoto Library/Masters/2012/09/29/20120929-235343/IMG_0407.MOV

2) 2 items
    /Users/caglar/Pictures/iPhoto Library/Masters/2012/07/18/20120718-032030/IMG_0411.MOV
    /Users/caglar/Pictures/iPhoto Library/Masters/2012/09/29/20120929-235343/IMG_0411.MOV

3) 2 items
    /Users/caglar/Pictures/iPhoto Library/Masters/2012/09/26/20120926-140126/IMG_0601.JPG
    /Users/caglar/Pictures/iPhoto Library/Masters/2012/09/29/20120929-235343/IMG_0780.JPG

4) 2 items
    /Users/caglar/Pictures/iPhoto Library/Masters/2012/07/18/20120718-032030/IMG_0420.MOV
    /Users/caglar/Pictures/iPhoto Library/Masters/2012/09/29/20120929-235525/IMG_0420.MOV

5) 2 items
    /Users/caglar/Pictures/iPhoto Library/Masters/2012/07/18/20120718-032030/IMG_0393.MOV
    /Users/caglar/Pictures/iPhoto Library/Masters/2012/09/29/20120929-235343/IMG_0393.MOV

6) 2 items
    /Users/caglar/Pictures/iPhoto Library/Masters/2012/07/18/20120718-032030/IMG_0413.MOV
    /Users/caglar/Pictures/iPhoto Library/Masters/2012/09/29/20120929-235343/IMG_0413.MOV

Checked 18859 files, found 6 different duplicated set of files (took 3.21 sec)
```

* Given extension:

```
[caglar@Zangetsu] ./dup -extension="jpg" /Users/caglar/Pictures/iPhoto\ Library/Masters/
Started to walk over /Users/caglar/Pictures/iPhoto Library/Masters and its sub-directories to find duplicated files

1) 2 items
    /Users/caglar/Pictures/iPhoto Library/Masters/2012/09/26/20120926-140126/IMG_0601.JPG
    /Users/caglar/Pictures/iPhoto Library/Masters/2012/09/29/20120929-235343/IMG_0780.JPG

Checked 18707 files, found 1 different duplicated set of files (took 1.66 sec)
```

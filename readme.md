# splot

_Simple command line plots_

`splot` reads values from stdin (one value per line) and creates simple plots for you.

### Usage

```
NAME:
   splot - simple plots

USAGE:
   splot [options] [arguments ...]

VERSION:
   1.0.0

OPTIONS:
   -w, --width value    width of the plot (default: 100)
   -h, --height value   height of the plot (an extra row will be added to this value) (default: 10)
   -m, --marker value   the character to use in the plot (default: .)
   -l, --line           when present, displays a line plot instead of full bars
       --help           print this usage
       --version        print version information
```

### Example

```
cody@redqueen:~/speedtest$ awk -F ", " 'NR > 500 {print $4}' speedtest.csv | splot
  25
  23                                         .
  21                     .             .     .                           .                       .
  19                     .             .     .   .     .   .   .         .                       .
  17           .         .         . . .     . . .   . .   .   .         .         .     .       .
  15   . . .   . . . . . .   . .   . . .   . . . . . . . . . . . .       . .   . . . . . . . .   .
  13 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
  11 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
   9 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
   7 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
```
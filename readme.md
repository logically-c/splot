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
cody@redqueen:~/speedtest$ awk -F ", " 'NR > 590 {print $4}' speedtest.csv | splot
27.30                                                                           .
25.98                                                                           .
24.66                                                                           .
23.34                                                             .             .
22.01               .             .   .             .             .             .
20.69 .             .             .   .             .             .     .       .     .
19.37 .       .     .             .   .         .   .         .   .     .       .     .
18.05 .       .     .             .   .         .   .         .   .     .       .     .
16.73 .       .     .       .     .   .         .   .       . . . .     .       .     .
15.40 .   .   .     . . . . . . . .   . . . .   . . .     . . . . . . . . . . . .     .
14.08 . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . .
```
# Splot

_Simple command line plots_

`splot` reads values from stdin (one value per line) and creates simple plots for you.

### Usage

```
NAME:
   splot - Simple plots

USAGE:
   splot [options] [arguments ...]

VERSION:
   1.0.0

OPTIONS:
   -w, --width value     (default: 100)
   -h, --height value    (default: 10)
   -m, --marker value    (default: .)
   -l, --line
       --help           print this usage
       --version        print version information
```

### Example

```bash
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
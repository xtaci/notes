#!/usr/bin/perl
# joke2.plx
use warnings;
use strict;

my @questions = qw(Java Python Perl C);
my @punchlines = (
	"None. Change it once, and it's the same everywhere.",
	"One. He just stands below the socket and the world revolves around him.",
	"A million. One to change it, the rest to try and do it in fewer lines.",
	'"CHANGE?!!"'
);

print "count: $#questions \n";

for (0..$#questions) {
	print "How many $questions[$_] ";
	print "programmers does it take to change a lightbulb?\n";
	sleep 2;
	print $punchlines[$_], "\n\n";
	sleep 1;
}

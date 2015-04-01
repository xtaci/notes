#!/usr/bin/perl
@ARGV = qw# larry moe curly #; # force these three files to be read
while (<>) {
	chomp;
	print "It was $_ that I saw in some stooge-like file!\n";
}

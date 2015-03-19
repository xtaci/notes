#!/usr/bin/perl
use warnings;
use strict;
my $w = "hello";
print qq/one $w two three four/;
print qw(one two three four);
print qw<one two three four>;
print qw{one two three four};
print qw[one two three four];
print qw|one two three four|;

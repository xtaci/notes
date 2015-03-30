#!/usr/bin/perl
use 5.010;
sub marine {
	state $n = 0; # private, persistent variable $n
	$n += 1;
	print "Hello, sailor number $n!\n";
}

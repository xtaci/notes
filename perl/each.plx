#!/usr/bin/perl
use 5.010;

@rocks = qw/ bedrock slate rubble granite /;
while( my( $index, $value ) = each @rocks ) {
	say "$index: $value";
}

@rocks = qw/ bedrock slate rubble granite /;
foreach $index ( 0 .. $#rocks ) {
	print "$index: $rocks[$index]\n";
}

<?php

require_once('Bench.php');


/*
* F(0) = 0
* F(1) = 1
* F(n) = F(n-1) + F(n-2) 
*/


function fib1 ($n) {
	if ($n <= 1) {
		return $n;
	}
	return fib1($n-1) + fib1($n-2);  
}


//echo fib1(11) . PHP_EOL;

$bench = new Bench(); 

$bench->run();
echo fib1(38) . PHP_EOL;
echo "fib1(38) time:" . $bench->stop()->result() . PHP_EOL;


function fib2($n) {
        if($n < 2){
        	return $n;
        }
        return fib2_sub(1, 1, $n);
}

function fib2_sub($a, $b, $c){
	if ($c <= 2){ 
		return $a; 
	}
    return fib2_sub($a+$b, $a, $c-1);
}

$bench = new Bench(); 
$bench->run();
echo fib2(38) . PHP_EOL;
echo "fib2(38) time:" . $bench->stop()->result() . PHP_EOL;












	

?>
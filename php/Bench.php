<?php

class Bench
{
    private $start;
    private $stop;
    private $result;
 
    public function run()
    {
        $this->start = microtime(true) * 1000;
    }
 
    public function stop()
    {
        $this->stop = microtime(true) * 1000;
        return $this;
    }

    public function result()
    {
        $this->result = $this->stop - $this->start;
        return round($this->result/1000, 2);
    }
 
    public function display()
    {
        echo sprintf("%.12f", $this->result()) . PHP_EOL;
    }
}

?>
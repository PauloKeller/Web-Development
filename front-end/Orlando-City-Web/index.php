<?php

require 'inc/Slim-2.x/Slim/Slim.php';

\Slim\Slim::registerAutoloader();


$app = new \Slim\Slim();



// GET route
$app->get(
    '/',
    function () {

            require_once("view/index.php");

    }
);

// GET route
$app->get(
    '/videos',
    function () {

            require_once("view/videos.php");

    }
);

$app->get('/produtos', function () {

            $sql = new Sql();

            $data = $sql->select("SELECT * FROM");

            json_encode($data);

    }
);




$app->run();

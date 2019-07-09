<?php include('config.php');

    //if user is not logged in, they cannot acess this page
    if (empty($_SESSION['username'])) {
        header('location: login.php');
    }
?>
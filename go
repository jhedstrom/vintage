<?php

// Github user or organization.
define('VINTAGE_USER', 'jhedstrom');

// Root.
define('VINTAGE_ROOT', getcwd());

// Data directory.
define('VINTAGE_DATA_DIRECTORY', __DIR__ . '/data');

require_once 'vendor/autoload.php';

$client = new Github\Client();
$repositories = $client->api('user')->repositories(VINTAGE_USER);

use Symfony\Component\Finder\Finder;

$make_files = array();
foreach ($repositories as $repo) {

  $name = $repo['name'];
  $clone_url = $repo['clone_url'];
  $directory = VINTAGE_DATA_DIRECTORY . '/' . $name;

  if (!file_exists($directory)) {
    chdir(VINTAGE_DATA_DIRECTORY);
    `git clone $clone_url`;
  }
  else {
    chdir($directory);
    `git pull`;
  }

  if ($makes = find_make_files($directory)) {
    $make_files[$name] = $makes;
  }

  // Go back.
  chdir(VINTAGE_ROOT);
}

var_dump($make_files);

/**
 * Find any makefiles in a given directory.
 */
function find_make_files($path) {
  $finder = new Finder();
  $iterator = $finder
    ->files()
    ->name('*.make')
    ->in($path);

  $files = array();
  foreach ($iterator as $found) {
    $files[$found->getRealPath()] = $found->getFileName();
  }

  return $files;
}

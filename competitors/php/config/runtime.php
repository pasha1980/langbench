<?php

use Runtime\Swoole\Runtime;

if(($_ENV['APP_RUNTIME'] ?? '') == Runtime::class) {
    $_SERVER['APP_RUNTIME_OPTIONS'] = [
        'host'     => '0.0.0.0',
        'port'     => '8080',
        'mode'     => SWOOLE_PROCESS,
        'settings' => [
            'daemonize' => false,
            'max_request' => 500,
            'package_max_length' => 10 << 20, //10mb
            'reactor_num' => 16,
            'send_yield' => true,
            'socket_buffer_size' => 10 << 20, //10mb
            'worker_num' => 16,
        ],
    ];
}

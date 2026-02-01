<?php

namespace App\Controller;

use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

readonly class CalcController
{
    #[Route(path: '/api/calc', methods: ['POST'])]
    public static function calc(Request $request): Response {
       $body = json_decode(
           json: $request->getContent(),
           depth: 5,
           flags: JSON_PARTIAL_OUTPUT_ON_ERROR,
       );
       $response = [
           'request_id' => $body->request_id,
           'stats' => [],
       ];
       $ids = [];

       foreach ($body->items as $itemValue) {
           if (isset($ids[$itemValue->id])) {
               continue;
           }

           $ids[$itemValue->id] = 0;

           foreach ($itemValue->tags as $tag) {

               if (!isset($response['stats'][$tag])) {
                   $response['stats'][$tag] = [
                       'sum' => $itemValue->value,
                       'count' => 1,
                   ];
                   continue;
               }

               $response['stats'][$tag]['sum'] += $itemValue->value;
               $response['stats'][$tag]['count']++;
            }
       }

       $response['checksum'] = hash('sha256', json_encode($response, JSON_PARTIAL_OUTPUT_ON_ERROR));
       return new Response(json_encode($response, JSON_PARTIAL_OUTPUT_ON_ERROR));
    }
}

<?php

declare(strict_types=1);

namespace App\MedCard\Controller;

use App\MedCard\Document\MedCard;
use Doctrine\ODM\MongoDB\DocumentManager;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

class GetMedCard
{
    public function __construct(
        private DocumentManager $dm
    ) {
    }

    #[Route(path: '/get/{id}', methods: ['GET'])]
    public function __invoke(string $id): JsonResponse
    {
        $repository = $this->dm->getRepository(MedCard::class);
        $card = $repository->find($id);

        if (!$card) {
            return new JsonResponse([
                'error' => 'Card not found'
            ], Response::HTTP_NOT_FOUND);
        }

        return new JsonResponse([
            'id' => $card->getId(),
            'vaccinations' => $card->getVaccination(),
            'allergies' => $card->getAllergy(),
            'diseases' => $card->getDisease(),
        ], Response::HTTP_OK);
    }
}

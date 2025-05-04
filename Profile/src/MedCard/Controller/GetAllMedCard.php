<?php

declare(strict_types=1);

namespace App\MedCard\Controller;

use App\MedCard\Document\MedCard;
use Doctrine\ODM\MongoDB\DocumentManager;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

class GetAllMedCard
{
    public function __construct(
        private DocumentManager $dm
    ) {
    }

    #[Route(path: '/get', methods: ['GET'])]
    public function __invoke(): JsonResponse
    {
        $repository = $this->dm->getRepository(MedCard::class);
        $cards = $repository->findAll();

        if (!$cards) {
            return new JsonResponse([
                'error' => 'Cards not found'
            ], Response::HTTP_NOT_FOUND);
        }

        $responceData = [];
        foreach ($cards as $card) {
            $data['id'] = $card->getId();
            $data['vaccinations'] = $card->getVaccination();
            $data['allergies'] = $card->getAllergy();
            $data['diseases'] = $card->getDisease();

            $responceData[] = $data;
        }

        return new JsonResponse(
            $responceData,
            Response::HTTP_OK,
        );
    }
}

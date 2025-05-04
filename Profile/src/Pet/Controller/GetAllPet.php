<?php

declare(strict_types=1);

namespace App\Pet\Controller;

use App\Pet\Document\Pet;
use Doctrine\ODM\MongoDB\DocumentManager;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

class GetAllPet
{
    public function __construct(
        private DocumentManager $dm
    ) {
    }

    #[Route(path: '/get', methods: ['GET'])]
    public function __invoke(): JsonResponse
    {
        $repository = $this->dm->getRepository(Pet::class);
        $pets = $repository->findAll();

        if (!$pets) {
            return new JsonResponse([
                'error' => 'Pets not found'
            ], Response::HTTP_NOT_FOUND);
        }

        $responceData = [];
        foreach ($pets as $pet) {
            $responceData[] = [
                'id' => $pet->getId(),
                'type' => $pet->getType(),
                'age' => $pet->getAge(),
                'breed' => $pet->getBreed(),
            ];
        }

        return new JsonResponse(
            $responceData,
            Response::HTTP_OK,
        );
    }
}

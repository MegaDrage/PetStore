<?php

declare(strict_types=1);

namespace App\Pet\Controller;

use App\Pet\Document\Pet;
use Doctrine\ODM\MongoDB\DocumentManager;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

class GetPet
{
    public function __construct(
        private DocumentManager $dm
    ) {
    }

    #[Route(path: '/get/{id}', methods: ['GET'])]
    public function __invoke(string $id): JsonResponse
    {
        $repository = $this->dm->getRepository(Pet::class);
        $pet = $repository->find($id);

        if (!$pet) {
            return new JsonResponse([
                'error' => 'Pet not found'
            ], Response::HTTP_NOT_FOUND);
        }

        return new JsonResponse([
            'id' => $pet->getId(),
            'type' => $pet->getType(),
            'age' => $pet->getAge(),
            'breed' => $pet->getBreed(),
        ], Response::HTTP_OK);
    }
}

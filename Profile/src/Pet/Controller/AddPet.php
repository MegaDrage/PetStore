<?php

declare(strict_types=1);

namespace App\Pet\Controller;

use App\Pet\Document\Pet;
use App\Pet\DTO\PetDTO;
use Doctrine\ODM\MongoDB\DocumentManager;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

class AddPet
{
    public function __construct(
        private DocumentManager $dm
    ) {
    }

    #[Route(path: '/add', methods: ['POST'])]
    public function __invoke(PetDTO $requestDTO): JsonResponse
    {
        $pet = new Pet(
            $requestDTO->getType(),
            $requestDTO->getAge(),
            $requestDTO->getBreed(),
        );

        $this->dm->persist($pet);
        $this->dm->flush();

        return new JsonResponse([
            'id' => $pet->getId(),
        ], Response::HTTP_OK);
    }
}

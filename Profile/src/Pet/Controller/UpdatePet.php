<?php

declare(strict_types=1);

namespace App\Pet\Controller;

use App\Pet\Document\Pet;
use App\Pet\DTO\PetDTO;
use Doctrine\ODM\MongoDB\DocumentManager;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

class UpdatePet
{
    public function __construct(
        private DocumentManager $dm
    ) {
    }

    #[Route(path: '/update/{id}', methods: ['PATCH'])]
    public function __invoke(string $id, PetDTO $requestDTO): JsonResponse
    {
        $repository = $this->dm->getRepository(Pet::class);
        $pet = $repository->find($id);

        if (!$pet) {
            return new JsonResponse([
                'error' => 'Pet not found',
            ], Response::HTTP_NOT_FOUND);
        }

        if ($requestDTO->getType() !== null) {
            $pet->setType($requestDTO->getType());
        }

        if ($requestDTO->getAge() !== null) {
            $pet->setAge($requestDTO->getAge());
        }

        if ($requestDTO->getBreed() !== null) {
            $pet->setBreed($requestDTO->getBreed());
        }

        $this->dm->persist($pet);
        $this->dm->flush();

        return new JsonResponse(
            null,
            Response::HTTP_OK,
        );
    }
}

<?php

declare(strict_types=1);

namespace App\Pet\Controller;

use App\Pet\Document\Pet;
use Doctrine\ODM\MongoDB\DocumentManager;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

class DeletePet
{
    public function __construct(
        private DocumentManager $dm
    ) {
    }

    #[Route(path: '/delete/{id}', methods: ['DELETE'])]
    public function __invoke(string $id): JsonResponse
    {
        $repository = $this->dm->getRepository(Pet::class);
        $pet = $repository->find($id);

        if (!$pet) {
            return new JsonResponse([
                'error' => 'Pet not found'
            ], Response::HTTP_NOT_FOUND);
        }

        $this->dm->remove($pet);
        $this->dm->flush();

        return new JsonResponse(
            null,
            Response::HTTP_OK,
        );
    }
}

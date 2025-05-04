<?php

declare(strict_types=1);

namespace App\MedCard\Controller;

use App\MedCard\Document\MedCard;
use App\Pet\Document\Pet;
use Doctrine\ODM\MongoDB\DocumentManager;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

class DeleteMedCard
{
    public function __construct(
        private DocumentManager $dm
    ) {
    }

    #[Route(path: '/delete/{id}', methods: ['DELETE'])]
    public function __invoke(string $id): JsonResponse
    {
        $repository = $this->dm->getRepository(MedCard::class);
        $card = $repository->find($id);

        if (!$card) {
            return new JsonResponse([
                'error' => 'Card not found'
            ], Response::HTTP_NOT_FOUND);
        }

        $pet = $this->dm->getRepository(Pet::class)->findOneBy(['cardId' => $card->getId()]);
        if ($pet) {
            $pet->setCardId(null);
            $this->dm->persist($pet);
        }

        $this->dm->remove($card);
        $this->dm->flush();

        return new JsonResponse(
            null,
            Response::HTTP_OK,
        );
    }
}

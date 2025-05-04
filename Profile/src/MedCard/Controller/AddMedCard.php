<?php

declare(strict_types=1);

namespace App\MedCard\Controller;

use App\MedCard\Document\Allergy;
use App\MedCard\Document\Disease;
use App\MedCard\Document\Vaccination;
use App\MedCard\DTO\MedCardDTO;
use App\MedCard\Document\MedCard;
use App\Pet\Document\Pet;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\ODM\MongoDB\DocumentManager;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

class AddMedCard
{
    public function __construct(
        private DocumentManager $dm
    ) {
    }

    #[Route(path: '/add/{petId}', methods: ['POST'])]
    public function __invoke(string $petId, MedCardDTO $requestDTO): JsonResponse
    {
        $repository = $this->dm->getRepository(Pet::class);
        $pet = $repository->find($petId);

        if (!$pet) {
            return new JsonResponse([
                'error' => 'Pet not found'
            ], Response::HTTP_NOT_FOUND);
        }

        if ($pet->getCardId() !== null) {
            return new JsonResponse([
                'error' => 'Card already exists'
            ], Response::HTTP_NOT_FOUND);
        }

        $allergies = [];
        foreach ($requestDTO->getAllergy() as $allergy) {
            $allergies[] = (new Allergy())->setName($allergy->getName());
        }
        $vaccinations = [];
        foreach ($requestDTO->getVaccination() as $vaccine) {
            $vaccinations[] = (new Vaccination())
                ->setName($vaccine->getName())
                ->setDate($vaccine->getDate());
        }
        $diseases = [];
        foreach ($requestDTO->getDisease() as $diesease) {
            $diseases[] = (new Disease())->setName($diesease->getName());
        }

        $card = new MedCard(
            new ArrayCollection($vaccinations),
            new ArrayCollection($allergies),
            new ArrayCollection($diseases),
        );

        $this->dm->persist($card);

        $pet->setCardId($card->getId());
        $this->dm->persist($pet);

        $this->dm->flush();

        return new JsonResponse([
            'id' => $card->getId(),
        ], Response::HTTP_OK);
    }
}

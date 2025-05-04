<?php

declare(strict_types=1);

namespace App\MedCard\Controller;

use App\MedCard\Document\Allergy;
use App\MedCard\Document\Disease;
use App\MedCard\Document\MedCard;
use App\MedCard\Document\Vaccination;
use App\MedCard\DTO\MedCardDTO;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\ODM\MongoDB\DocumentManager;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

class UpdateMedCard
{
    public function __construct(
        private DocumentManager $dm
    ) {
    }

    #[Route(path: '/update/{id}', methods: ['PATCH'])]
    public function __invoke(string $id, MedCardDTO $requestDTO): JsonResponse
    {
        $repository = $this->dm->getRepository(MedCard::class);
        $card = $repository->find($id);

        if (!$card) {
            return new JsonResponse([
                'error' => 'Card not found',
            ], Response::HTTP_NOT_FOUND);
        }

        $allergies = [];
        foreach ($requestDTO->getAllergy()??[] as $allergy) {
            $allergies[] = (new Allergy())->setName($allergy->getName());
        }
        if (!empty($allergies)) {
            $card->setAllergy(new ArrayCollection($allergies));
        }

        $vaccinations = [];
        foreach ($requestDTO->getVaccination()??[] as $vaccine) {
            $vaccinations[] = (new Vaccination())
                ->setName($vaccine->getName())
                ->setDate($vaccine->getDate());
        }
        if (!empty($vaccinations)) {
            $card->setVaccination(new ArrayCollection($vaccinations));
        }

        $diseases = [];
        foreach ($requestDTO->getDisease()??[] as $diesease) {
            $diseases[] = (new Disease())->setName($diesease->getName());
        }
        if (!empty($diseases)) {
            $card->setDisease(new ArrayCollection($diseases));
        }

        $this->dm->persist($card);
        $this->dm->flush();

        return new JsonResponse(
            null,
            Response::HTTP_OK,
        );
    }
}

<?php

declare(strict_types=1);

namespace App\MedCard\DTO;

use App\RequestDtoArgumentInterface;
use JMS\Serializer\Annotation as JMS;

class MedCardDTO implements RequestDtoArgumentInterface
{
    #[JMS\Type('array<App\MedCard\DTO\VaccinationDTO>')]
    #[JMS\SerializedName('vaccination')]
    private ?array $vaccination = null;

    #[JMS\Type('array<App\MedCard\DTO\AllergyDTO>')]
    #[JMS\SerializedName('allergy')]
    private ?array $allergy = null;

    #[JMS\Type('array<App\MedCard\DTO\DiseaseDTO>')]
    #[JMS\SerializedName('disease')]
    private ?array $disease = null;

    public function getVaccination(): ?array
    {
        return $this->vaccination;
    }

    public function getAllergy(): ?array
    {
        return $this->allergy;
    }

    public function getDisease(): ?array
    {
        return $this->disease;
    }
}

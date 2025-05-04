<?php

declare(strict_types=1);

namespace App\Pet\DTO;

use App\RequestDtoArgumentInterface;
use JMS\Serializer\Annotation as JMS;

class PetDTO implements RequestDtoArgumentInterface
{
    #[JMS\Type('string')]
    #[JMS\SerializedName('type')]
    private ?string $type = null;

    #[JMS\Type('int')]
    #[JMS\SerializedName('age')]
    private ?int $age = null;

    #[JMS\Type('string')]
    #[JMS\SerializedName('breed')]
    private ?string $breed = null;

    public function getType(): ?string
    {
        return $this->type;
    }

    public function getAge(): ?int
    {
        return $this->age;
    }

    public function getBreed(): ?string
    {
        return $this->breed;
    }
}

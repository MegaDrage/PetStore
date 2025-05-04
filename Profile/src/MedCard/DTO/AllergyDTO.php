<?php

declare(strict_types=1);

namespace App\MedCard\DTO;

use App\RequestDtoArgumentInterface;
use JMS\Serializer\Annotation as JMS;

class AllergyDTO implements RequestDtoArgumentInterface
{
    #[JMS\Type('string')]
    #[JMS\SerializedName('name')]
    private string $name;

    public function getName(): ?string
    {
        return $this->name;
    }

    public function setName(string $name): self
    {
        $this->name = $name;
        return $this;
    }
}

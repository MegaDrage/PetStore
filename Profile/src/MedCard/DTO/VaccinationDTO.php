<?php

declare(strict_types=1);

namespace App\MedCard\DTO;

use App\RequestDtoArgumentInterface;
use DateTimeInterface;
use JMS\Serializer\Annotation as JMS;

class VaccinationDTO implements RequestDtoArgumentInterface
{
    #[JMS\Type('string')]
    #[JMS\SerializedName('name')]
    private string $name;

    #[JMS\Type("DateTime<'Y-m-d'>")]
    #[JMS\SerializedName('date')]
    private ?DateTimeInterface $date = null;

    public function getName(): string
    {
        return $this->name;
    }

    public function setName(string $name): self
    {
        $this->name = $name;
        return $this;
    }

    public function getDate(): ?DateTimeInterface
    {
        return $this->date;
    }

    public function setDate(?DateTimeInterface $date): self
    {
        $this->date = $date;
        return $this;
    }
}

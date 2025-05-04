<?php

declare(strict_types=1);

namespace App\MedCard\Document;

use Doctrine\ODM\MongoDB\Mapping\Annotations as MongoDB;
use DateTimeInterface;

#[MongoDB\EmbeddedDocument]
class Vaccination
{
    #[MongoDB\Field(type: 'string')]
    private string $name;

    #[MongoDB\Field(type: 'date')]
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

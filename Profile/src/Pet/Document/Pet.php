<?php

declare(strict_types=1);

namespace App\Pet\Document;

use App\MedCard\Document\MedCard;
use Doctrine\ODM\MongoDB\Mapping\Annotations as MongoDB;

#[MongoDB\Document]
class Pet
{
    #[MongoDB\Id]
    private string $id;

    #[MongoDB\Field(type: 'string')]
    private ?string $type;

    #[MongoDB\Field(type: 'int')]
    private ?int $age;

    #[MongoDB\Field(type: 'string')]
    private ?string $breed;

    #[MongoDB\Field(type: 'string')]
    private ?string $cardId = null;

    public function __construct(?string $type = null, ?int $age = null, ?string $breed = null)
    {
        $this->type = $type;
        $this->age = $age;
        $this->breed = $breed;
    }

    public function getId(): ?string
    {
        return $this->id;
    }

    public function getType(): string
    {
        return $this->type;
    }

    public function setType(string $type): self
    {
        $this->type = $type;
        return $this;
    }

    public function getAge(): int
    {
        return $this->age;
    }

    public function setAge(int $age): self
    {
        $this->age = $age;
        return $this;
    }

    public function getBreed(): string
    {
        return $this->breed;
    }

    public function setBreed(string $breed): self
    {
        $this->breed = $breed;
        return $this;
    }

    public function getCardId(): ?string
    {
        return $this->cardId;
    }

    public function setCardId(?string $id): self
    {
        $this->cardId = $id;
        return $this;
    }
}
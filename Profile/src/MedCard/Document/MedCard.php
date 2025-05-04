<?php

declare(strict_types=1);

namespace App\MedCard\Document;

use App\MedCard\Document\Allergy;
use App\MedCard\Document\Disease;
use App\MedCard\Document\Vaccination;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\ODM\MongoDB\Mapping\Annotations as MongoDB;

#[MongoDB\Document(collection: 'med_cards')]
class MedCard
{
    #[MongoDB\Id]
    private string $id;

    #[MongoDB\EmbedMany(targetDocument: Vaccination::class)]
    private ?ArrayCollection $vaccination = null;

    #[MongoDB\EmbedMany(targetDocument: Allergy::class)]
    private ?ArrayCollection $allergy = null;

    #[MongoDB\EmbedMany(targetDocument: Disease::class)]
    private ?ArrayCollection $disease = null;


    public function __construct(?ArrayCollection $vaccination, ?ArrayCollection $allergy, ?ArrayCollection $disease)
    {
        $this->vaccination = $vaccination;
        $this->allergy = $allergy;
        $this->disease = $disease;
    }

    public function getId(): string
    {
        return $this->id;
    }

    public function getVaccination(): ?array
    {
        if ($this->vaccination === null) {
            return null;
        }

        $result = [];
        foreach ($this->vaccination as $vaccination) {
            $data['name'] = $vaccination->getName();
            $data['date'] = $vaccination->getDate();
            $result[] = $data;
        }
        return $result;
    }

    public function setVaccination(?ArrayCollection $vaccination): self
    {
        $this->vaccination = $vaccination;
        return $this;
    }

    public function getAllergy(): ?array
    {
        if ($this->allergy === null) {
            return null;
        }

        $result = [];
        foreach ($this->allergy as $allergy) {
            $data['name'] = $allergy->getName();
            $result[] = $data;
        }
        return $result;
    }

    public function setAllergy(?ArrayCollection $allergy): self
    {
        $this->allergy = $allergy;
        return $this;
    }

    public function getDisease(): ?array
    {
        if ($this->disease === null) {
            return null;
        }

        $result = [];
        foreach ($this->disease as $disease) {
            $data['name'] = $disease->getName();
            $result[] = $data;
        }
        return $result;
    }

    public function setDisease(?ArrayCollection $disease): self
    {
        $this->disease = $disease;
        return $this;
    }
}
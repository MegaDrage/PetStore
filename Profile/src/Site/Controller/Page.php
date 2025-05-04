<?php

declare(strict_types=1);

namespace App\Pet\Controller;

use Symfony\Bundle\FrameworkBundle\Controller\AbstractController;
use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

class Page extends AbstractController
{
    #[Route(path: '/', methods: ['GET'])]
    public function index(): Response
    {
        return $this->render('index.html.twig');
    }

    #[Route(path: '/adminka', methods: ['GET'])]
    public function adminka(): Response
    {
        return $this->render('adminka.html.twig');
    }

    #[Route(path: '/analysis', methods: ['GET'])]
    public function analysis(): Response
    {
        return $this->render('analysis.html.twig');
    }

    #[Route(path: '/appointment', methods: ['GET'])]
    public function appointment(): Response
    {
        return $this->render('appointment.html.twig');
    }

    #[Route(path: '/base', methods: ['GET'])]
    public function base(): Response
    {
        return $this->render('base.html.twig');
    }

    #[Route(path: '/chat', methods: ['GET'])]
    public function chat(): Response
    {
        return $this->render('chat.html.twig');
    }

    #[Route(path: '/forum_list', methods: ['GET'])]
    public function forum_list(): Response
    {
        return $this->render('forum-list.html.twig');
    }

    #[Route(path: '/forum-topic1', methods: ['GET'])]
    public function forum_topic1(): Response
    {
        return $this->render('forum-topic1.html.twig');
    }

    #[Route(path: '/forum-topic2', methods: ['GET'])]
    public function forum_topic2(): Response
    {
        return $this->render('forum-topic2.html.twig');
    }

    #[Route(path: '/inventory', methods: ['GET'])]
    public function inventory(): Response
    {
        return $this->render('inventory.html.twig');
    }

    #[Route(path: '/my_pets', methods: ['GET'])]
    public function my_pets(): Response
    {
        return $this->render('my_pets.html.twig');
    }

    #[Route(path: '/pet_profil', methods: ['GET'])]
    public function pet_profil(): Response
    {
        return $this->render('pet_profil.html.twig');
    }

    #[Route(path: '/schedule', methods: ['GET'])]
    public function schedule(): Response
    {
        return $this->render('schedule.html.twig');
    }
}

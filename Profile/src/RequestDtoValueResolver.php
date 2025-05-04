<?php

namespace App;

use JMS\Serializer\SerializerInterface;
use Symfony\Component\HttpFoundation\Request;
use Symfony\Component\HttpKernel\Controller\ValueResolverInterface;
use Symfony\Component\HttpKernel\ControllerMetadata\ArgumentMetadata;

class RequestDtoValueResolver implements ValueResolverInterface
{
    public function __construct(
        private readonly SerializerInterface $serializer
    ) {}

    public function resolve(Request $request, ArgumentMetadata $argument): iterable
    {
        if (!is_a($argument->getType(), RequestDtoArgumentInterface::class, true)) {
            return [];
        }

        $data = $request->getContent() ?: '{}';

        yield $this->serializer->deserialize(
            $data,
            $argument->getType(),
            'json'
        );
    }
}
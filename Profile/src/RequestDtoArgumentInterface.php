<?php

declare(strict_types=1);

namespace App;

/**
 * Этот интерфейс должны реализовывать DTO с входными параметрами контроллера
 * Благодаря этому на них сработает нужный ArgumentResolver
 *
 * @see DtoResolver::supports()
 */
interface RequestDtoArgumentInterface
{
}

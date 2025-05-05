using Microsoft.AspNetCore.Mvc;
using TelehealthService.Api.Contracts;
using TelehealthService.Core.Abstractions;
using TelehealthService.Api.Resources;

namespace TelehealthService.Api.Controllers;

[ApiController]
[Route("[controller]")]
public class TelehealthController : ControllerBase
{
    private readonly IJitsiService _jitsiService;
    private readonly IAiAssistantService _aiAssistantService;

    public TelehealthController(
        IJitsiService jitsiService, 
        IAiAssistantService aiAssistantService)
    {
        _jitsiService = jitsiService;
        _aiAssistantService = aiAssistantService;
    }
    
    [HttpPost("start")]
    public IActionResult StartSession([FromBody] ConsultationRequest request)
    {
        var roomUrl = _jitsiService.GenerateRoomUrl(request.PetOwnerId);
        return Ok(new { RoomUrl = roomUrl });
    }
    
    [HttpPost("diagnose")]
    public async Task<IActionResult> Diagnose([FromBody] DiagnoseRequest request)
    {
        try
        {
            var prompt = string.Format(
                Prompts.VeterinaryPrompt,
                request.AnimalKind,
                string.Join(", ", request.Symptoms)
            );
            
            var diagnosis = await _aiAssistantService.GetResponseAsync(prompt);
            return Ok(new { Diagnosis = diagnosis });
        }
        catch (Exception ex)
        {
            return StatusCode(500, $"AiController Error: {ex.Message}");
        }
    }
}
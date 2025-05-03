namespace PetStore.JitsiAPI.Models;

public class ConsultationSession
{
    public string SessionId { get; set; } = Guid.NewGuid().ToString();
    public string VeterinarianId { get; set; }
    public string PetOwnerId { get; set; }
    public DateTime StartTime { get; set; }
}
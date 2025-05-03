namespace TelehealthService.Api.Contracts;

public class ConsultationRequest
{
    public string VetDoctorId { get; set; }
    public string PetOwnerId { get; set; }
    public DateTime DateTime { get; set; }
}
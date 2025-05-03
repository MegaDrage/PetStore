namespace TelehealthService.Api.Contracts;

public class DiagnoseRequest
{
    public string AnimalKind { get; set; }
    public List<string> Symptoms { get; set; }
}
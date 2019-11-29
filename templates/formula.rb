class {{.FormulaClassName}} < Formula
    desc "{{.Description}}"
    homepage "{{.Homepage}}"
    url "{{.BinaryURL}}"
    sha256 "{{.Sha256}}"
    version "{{.Version}}"

  def install
    mv Dir.glob("{{.Formula}}-*").first, "{{.Formula}}"
    bin.install "{{.Formula}}"
  end

  test do
    "{{.Formula}}" help
  end
end
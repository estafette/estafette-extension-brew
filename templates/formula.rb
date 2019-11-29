class {{.FormulaClassName}} < Formula
    desc "{{.Description}}"
    homepage "{{.Homepage}}"
    url "{{.BinaryURL}}"
    sha256 "{{.Sha256}}"
    version "{{.Version}}"

  def install
    bin.install "{{.Filename}}" => "{{.Formula}}"
  end

  test do
    system "{{.Formula}}", "help"
  end
end
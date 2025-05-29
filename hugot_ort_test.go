//go:build !NOORT || ALL

package hugot

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/prasanthmj/hugot/options"
	"github.com/prasanthmj/hugot/util"
)

// FEATURE EXTRACTION

func TestFeatureExtractionPipelineORT(t *testing.T) {
	opts := []options.WithOption{options.WithOnnxLibraryPath(onnxRuntimeSharedLibrary)}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	featureExtractionPipeline(t, session)
}

func TestFeatureExtractionPipelineORTCuda(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.SkipNow()
	}
	opts := []options.WithOption{
		options.WithOnnxLibraryPath("/usr/lib64/onnxruntime-gpu/libonnxruntime.so"),
		options.WithCuda(map[string]string{
			"device_id": "0",
		}),
	}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	featureExtractionPipeline(t, session)
}

func TestFeatureExtractionPipelineValidationORT(t *testing.T) {
	opts := []options.WithOption{options.WithOnnxLibraryPath(onnxRuntimeSharedLibrary)}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	featureExtractionPipelineValidation(t, session)
}

// Text classification

func TestTextClassificationPipelineORT(t *testing.T) {
	opts := []options.WithOption{
		options.WithOnnxLibraryPath(onnxRuntimeSharedLibrary),
		options.WithTelemetry(),
		options.WithCpuMemArena(true),
		options.WithMemPattern(true),
		options.WithIntraOpNumThreads(1),
		options.WithInterOpNumThreads(1),
	}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	textClassificationPipeline(t, session)
}

func TestTextClassificationPipelineORTCuda(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.SkipNow()
	}
	opts := []options.WithOption{
		options.WithOnnxLibraryPath("/usr/lib64/onnxruntime-gpu/libonnxruntime.so"),
		options.WithCuda(map[string]string{
			"device_id": "0",
		}),
	}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	textClassificationPipeline(t, session)
}

func TestTextClassificationPipelineMultiORT(t *testing.T) {
	opts := []options.WithOption{
		options.WithOnnxLibraryPath(onnxRuntimeSharedLibrary),
		options.WithTelemetry(),
		options.WithCpuMemArena(true),
		options.WithMemPattern(true),
		options.WithIntraOpNumThreads(1),
		options.WithInterOpNumThreads(1),
	}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	textClassificationPipelineMulti(t, session)
}

func TestTextClassificationPipelineORTMultiCuda(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.SkipNow()
	}
	opts := []options.WithOption{
		options.WithOnnxLibraryPath("/usr/lib64/onnxruntime-gpu/libonnxruntime.so"),
		options.WithCuda(map[string]string{
			"device_id": "0",
		}),
	}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	textClassificationPipelineMulti(t, session)
}

func TestTextClassificationPipelineValidationORT(t *testing.T) {
	opts := []options.WithOption{options.WithOnnxLibraryPath(onnxRuntimeSharedLibrary)}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	textClassificationPipelineValidation(t, session)
}

// Token classification

func TestTokenClassificationPipelineORT(t *testing.T) {
	opts := []options.WithOption{options.WithOnnxLibraryPath(onnxRuntimeSharedLibrary)}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	tokenClassificationPipeline(t, session)
}

func TestTokenClassificationPipelineORTCuda(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.SkipNow()
	}
	opts := []options.WithOption{
		options.WithOnnxLibraryPath("/usr/lib64/onnxruntime-gpu/libonnxruntime.so"),
		options.WithCuda(map[string]string{
			"device_id": "0",
		}),
	}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	tokenClassificationPipeline(t, session)
}

func TestTokenClassificationPipelineValidationORT(t *testing.T) {
	opts := []options.WithOption{options.WithOnnxLibraryPath(onnxRuntimeSharedLibrary)}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	tokenClassificationPipelineValidation(t, session)
}

// Zero shot

func TestZeroShotClassificationPipelineORT(t *testing.T) {
	opts := []options.WithOption{options.WithOnnxLibraryPath(onnxRuntimeSharedLibrary)}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	zeroShotClassificationPipeline(t, session)
}

func TestZeroShotClassificationPipelineORTCuda(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.SkipNow()
	}
	opts := []options.WithOption{
		options.WithOnnxLibraryPath("/usr/lib64/onnxruntime-gpu/libonnxruntime.so"),
		options.WithCuda(map[string]string{
			"device_id": "0",
		}),
	}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	zeroShotClassificationPipeline(t, session)
}

func TestZeroShotClassificationPipelineValidationORT(t *testing.T) {
	opts := []options.WithOption{options.WithOnnxLibraryPath(onnxRuntimeSharedLibrary)}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	zeroShotClassificationPipelineValidation(t, session)
}

// No Same Name

func TestNoSameNamePipelineORT(t *testing.T) {
	opts := []options.WithOption{options.WithOnnxLibraryPath(onnxRuntimeSharedLibrary)}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	noSameNamePipeline(t, session)
}

func TestClosePipelineORT(t *testing.T) {
	opts := []options.WithOption{options.WithOnnxLibraryPath(onnxRuntimeSharedLibrary)}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	destroyPipelines(t, session)
}

// Thread safety

func TestThreadSafetyORT(t *testing.T) {
	opts := []options.WithOption{options.WithOnnxLibraryPath(onnxRuntimeSharedLibrary)}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	threadSafety(t, session)
}

func TestThreadSafetyORTCuda(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.SkipNow()
	}
	opts := []options.WithOption{
		options.WithOnnxLibraryPath("/usr/lib64/onnxruntime-gpu/libonnxruntime.so"),
		options.WithCuda(map[string]string{
			"device_id": "0",
		}),
	}
	session, err := NewORTSession(opts...)
	check(t, err)
	defer func(session *Session) {
		destroyErr := session.Destroy()
		check(t, destroyErr)
	}(session)
	threadSafety(t, session)
}

// Benchmarks

func runBenchmarkEmbedding(strings *[]string, cuda bool) {
	var opts []options.WithOption
	switch cuda {
	case true:
		opts = []options.WithOption{
			options.WithOnnxLibraryPath("/usr/lib64/onnxruntime-gpu/libonnxruntime.so"),
			options.WithCuda(map[string]string{
				"device_id": "0",
			}),
		}
	default:
		opts = []options.WithOption{options.WithOnnxLibraryPath("/usr/lib64/onnxruntime.so")}
	}
	session, err := NewORTSession(opts...)
	if err != nil {
		panic(err)
	}

	defer func(session *Session) {
		errDestroy := session.Destroy()
		if errDestroy != nil {
			panic(errDestroy)
		}
	}(session)

	modelPath := "./models/sentence-transformers_all-MiniLM-L6-v2"
	config := FeatureExtractionConfig{
		ModelPath: modelPath,
		Name:      "benchmarkEmbedding",
	}
	pipelineEmbedder, err2 := NewPipeline(session, config)
	if err2 != nil {
		panic(err2)
	}
	res, err := pipelineEmbedder.Run(*strings)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(res.GetOutput()))
}

func BenchmarkCudaEmbedding(b *testing.B) {
	if os.Getenv("CI") != "" {
		b.SkipNow()
	}
	p := make([]string, 30000)
	for i := 0; i < 30000; i++ {
		p[i] = "The goal of this library is to provide an easy, scalable, and hassle-free way to run huggingface transformer pipelines in golang applications."
	}
	for i := 0; i < b.N; i++ {
		runBenchmarkEmbedding(&p, true)
	}
}

func BenchmarkCPUEmbedding(b *testing.B) {
	if os.Getenv("CI") != "" {
		b.SkipNow()
	}
	p := make([]string, 30000)
	for i := 0; i < 30000; i++ {
		p[i] = "The goal of this library is to provide an easy, scalable, and hassle-free way to run huggingface transformer pipelines in golang applications."
	}
	for i := 0; i < b.N; i++ {
		runBenchmarkEmbedding(&p, false)
	}
}

// README: test the readme examples

func TestReadmeExample(t *testing.T) {
	check := func(err error) {
		if err != nil {
			panic(err.Error())
		}
	}

	// start a new session. By default, this looks for the onnxruntime.so library in its default path, e.g. /usr/lib/onnxruntime.so
	// if your onnxruntime.so is somewhere else, you can explicitly set it by using WithOnnxLibraryPath
	session, err := NewORTSession(options.WithOnnxLibraryPath("/usr/lib64/onnxruntime.so"))
	check(err)
	// A successfully created hugot session needs to be destroyed when you're done
	defer func(session *Session) {
		err := session.Destroy()
		check(err)
	}(session)

	// Let's download an onnx sentiment test classification model in the current directory
	// note: if you compile your library with build flag NODOWNLOAD, this will exclude the downloader.
	// Useful in case you just want the core engine (because you already have the models) and want to
	// drop the dependency on huggingfaceModelDownloader.
	modelPath, err := DownloadModel("KnightsAnalytics/distilbert-base-uncased-finetuned-sst-2-english", "./", NewDownloadOptions())
	check(err)

	defer func(modelPath string) {
		err := util.DeleteFile(modelPath)
		if err != nil {
			t.FailNow()
		}
	}(modelPath)

	// we now create the configuration for the text classification pipeline we want to create.
	// Options to the pipeline can be set here using the Options field
	config := TextClassificationConfig{
		ModelPath: modelPath,
		Name:      "testPipeline",
	}
	// then we create out pipeline.
	// Note: the pipeline will also be added to the session object so all pipelines can be destroyed at once
	sentimentPipeline, err := NewPipeline(session, config)
	check(err)

	// we can now use the pipeline for prediction on a batch of strings
	batch := []string{"This movie is disgustingly good !", "The director tried too much"}
	batchResult, err := sentimentPipeline.RunPipeline(batch)
	check(err)

	// and do whatever we want with it :)
	s, err := json.Marshal(batchResult)
	check(err)
	fmt.Println(string(s))
	// {"ClassificationOutputs":[[{"Label":"POSITIVE","Score":0.9998536}],[{"Label":"NEGATIVE","Score":0.99752176}]]}
}
